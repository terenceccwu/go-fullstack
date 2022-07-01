import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  build: {
    outDir: 'web_app/static/dist',
    assetsDir: '.',
    rollupOptions: {
      input: 'web_app/static/src/index.jsx',
      output: {
        entryFileNames: `[name].js`,
        chunkFileNames: `[name].js`,
        assetFileNames: `[name].[ext]`
      }
    }
  },
  // resolve: {
  //   alias: {
  //   },
  // },
  server: {
    port: 3001
  }
});
