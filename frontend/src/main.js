import './style.css'
import './app.css'; // 👈 Tailwind styles

import App from './App.svelte'

const app = new App({
  target: document.getElementById('app')
})

export default app
