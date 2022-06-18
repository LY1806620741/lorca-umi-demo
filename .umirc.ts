import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  history: { type: 'hash' }, //单页路由问题
  routes: [
    { path: '/', component: '@/pages/index' },
    { path: '/about', component: '@/pages/about' },
  ],
  title: 'Demo',
  fastRefresh: {},
});
