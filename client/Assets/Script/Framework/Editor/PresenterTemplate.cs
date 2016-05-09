﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Framework
{
    class PresenterTemplate
    {
        public static string ClassName(DataContext ctx)
        {
            return ctx.Name + "Presenter";
        }



        // Presenter中的事件名
        public static string Event(DataContext ctx)
        {
            return "On" + ctx.Name + "Changed";
        }

        public static string Property( DataContext ctx )
        {
            return ctx.Name;
        }

        public static string PropertyType( DataContext ctx )
        {
            switch( ctx.Type )
            {
                case WidgetType.InputField:
                case WidgetType.Text:
                    return "string";
            }

            return "unknown";
        }

        public static string Command(DataContext ctx)
        {
            return "Cmd_" + ctx.Name;
        }

        static bool HasPropertyChangedNotify( DataContext propContext )
        {            
            return propContext.SyncType ==  DataSyncType.PresenterToView || propContext.SyncType == DataSyncType.TwoWay;          
        }

        public static void PropertyBody(CodeGenerator gen, DataContext presenterContext, DataContext propContext )
        {
            if ( propContext.Type == WidgetType.ScrollRect )
            {
                gen.PrintLine("public Framework.ObservableCollection<int, ", ClassName(presenterContext), ">", Property(propContext), "{ get; set; }");
            }
            else
            {
                if (HasPropertyChangedNotify(propContext))
                {
                    gen.PrintLine("public Action ", Event(propContext), ";");
                }

                gen.PrintLine("public ", PropertyType(propContext), " ", Property(propContext));
                gen.PrintLine("{");
                gen.In();

                gen.PrintLine("get");
                gen.PrintLine("{");
                gen.In();
                gen.PrintLine("return _Model.", Property(propContext));
                gen.Out();
                gen.PrintLine("}"); // get


                gen.PrintLine("set");
                gen.PrintLine("{");
                gen.In();
                gen.PrintLine("_Model.", Property(propContext), " = value;");

                gen.PrintLine();

                // 有Presenter同步到View时, 调用通知
                if (HasPropertyChangedNotify(propContext))
                {
                    gen.PrintLine("if ( ", Event(propContext), " != null )");
                    gen.PrintLine("{");
                    gen.In();
                    gen.PrintLine(Event(propContext), "();");
                    gen.Out();
                    gen.PrintLine("}"); // set
                }


                gen.Out();
                gen.PrintLine("}"); // set


                gen.Out();
                gen.PrintLine("}"); // Property
            }
        }

        public static void PropertyInit(CodeGenerator gen, DataContext presenterContext, DataContext propContext )
        {
            if ( propContext.Type == WidgetType.ScrollRect )
            {
                gen.PrintLine(Property(propContext), " = new Framework.ObservableCollection<int, ", ClassName(presenterContext), ">();");
            }
        }

        public static void Class(CodeGenerator gen, DataContext presenterContext, List<DataContext> propContextList)
        {
            gen.PrintLine("// Generated by github.com/davyxu/cellorigin");
            gen.PrintLine("using UnityEngine;");
            gen.PrintLine("using UnityEngine.UI;");
            gen.PrintLine();

            gen.PrintLine("partial class ", ClassName(presenterContext), " : Framework.BasePresenter");
            gen.PrintLine("{");
            gen.In();

            gen.PrintLine(ModelTemplate.ClassName(presenterContext), " _Model;");
            gen.PrintLine();

            // TODO 网络Peer依赖

            foreach( DataContext propContext in propContextList )
            {
                PropertyBody(gen, presenterContext, propContext);
            }

            gen.PrintLine("public override void Init( )");
            gen.PrintLine("{");
            gen.In();

            // TODO Model初始化多种方法

            // TODO Collection初始化

            foreach (DataContext propContext in propContextList)
            {
                PropertyInit(gen, presenterContext, propContext);
            }

            // TODO 网络获取及消息绑定

            gen.Out();
            gen.PrintLine("}"); // Bind
            gen.PrintLine();

            gen.Out();
            gen.PrintLine("}"); // Class
        }
    }
}