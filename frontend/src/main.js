import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import './style/common.scss'

const app = createApp(App);
app.use(router).mount('#app')
