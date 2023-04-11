import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import './assets/main.css'
import './common/protocol.js'
import './common/starx.js'
import 'element-plus/dist/index.css'
import './assets/icon/iconfont.css'
var app = createApp(App)
app.use(ElementPlus).mount('#app')
