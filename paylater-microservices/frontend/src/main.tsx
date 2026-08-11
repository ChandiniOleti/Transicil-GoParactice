import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux'

import { store } from './app/store'
import { setupApiInterceptors } from './services/setupApiInterceptors'
import { ThemeProvider } from './context/ThemeContext'
import './index.css'
import './components/components.css'
import App from './App.tsx'

setupApiInterceptors(store)

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider>
      <Provider store={store}>
        <App />
      </Provider>
    </ThemeProvider>
  </StrictMode>,
)
