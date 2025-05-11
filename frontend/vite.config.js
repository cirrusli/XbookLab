import { defineConfig } from 'vite'
import { resolve } from 'path'
import vue from '@vitejs/plugin-vue'
import { viteMockServe } from 'vite-plugin-mock'

import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    viteMockServe({
      mockPath: 'mock',
      localEnabled: true,
    }),
    Components({
      resolvers: [
        AntDesignVueResolver({
          importStyle: false, // css in js
        }),
      ],
    }),
  ],

  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "${resolve(__dirname, 'src/style/common.scss')}";`
      }
    }
  },

})
