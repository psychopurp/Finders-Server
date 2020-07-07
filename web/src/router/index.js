import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);
import Layout from "@/layout";

/**
 * 路由相关属性说明
 * hidden: 当设置hidden为true时，意思不在sideBars侧边栏中显示
 * mete{
 * title: xxx,  设置sideBars侧边栏名称
 * icon: xxx,  设置sideBars侧边栏图标
 * noCache: true  当设置为true时不缓存该路由页面
 * }
 */

/*通用routers 不需要鉴权*/
export const currencyRoutes = [
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/login"),
    meta: { title: "登录页" },
    hidden: true,
  },
  {
    path: "/404",
    name: "404",
    component: () => import("@/views/error-page/404.vue"),
    hidden: true,
  },

  {
    path: "/",
    name: "Home",
    component: Layout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        component: () => import("@/views/dashboard"),
        meta: { title: "首页", icon: "el-icon-s-data" },
      },
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  routes: currencyRoutes,
});

export default router;
