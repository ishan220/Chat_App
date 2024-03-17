import React from 'react';
import ReactDOM from 'react-dom/client';
import {createBrowserRouter,RouterProvider} from 'react-router-dom'

import './index.css';
import App from './App';
import Register from './components/Register.js'
import LoginUser from './components/LoginUser'
import Landing from './components/Landing.js'
import Chat from './components/Chat'


import reportWebVitals from './reportWebVitals';
const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
    children: [
      {
        path: "/",
        element: <Landing/>,
      },
      {
        path: "/login",
        element: <LoginUser/>,
      },
      {
        path: "/register",
        element: <Register/>,
      },
      {
        path: "/chat",
        element: <Chat/>,
      },
      
    ],
 },])
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(<RouterProvider router={router}/>);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
