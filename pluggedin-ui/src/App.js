import './App.css';
import {ChakraProvider} from '@chakra-ui/react'
import Chat from './chat/Chat';
import { Route, Navigate, Routes, useLocation, BrowserRouter } from 'react-router-dom';
import { AuthProvider, useAuth } from './auth/UseAuth';
//import { lazy, Suspense } from 'react';
import Login from './login/Login';

// const AsyncRoute = ({element, ...props}) => {
//   console.log("async")
//   return <Route {...props} element={element}> </Route>
// }

const AuthenticatedRoute = ({children}) => {
  let { user } = useAuth()
  let location = useLocation() 
  
  if(!user) {
    return <Navigate to="/login" state={{ from: location }}/>;
  }

  return children
}

function App() {
  
  return (
    <BrowserRouter>
      <ChakraProvider>
        <AuthProvider>
          <Routes>
              <Route 
              path="/"
              element={
                <AuthenticatedRoute>
                  <Chat />
                </AuthenticatedRoute>
              } />
              <Route 
                exact
                path="/login"
                element={
                  <Login />
                }
              />
          </Routes>
        </AuthProvider>
      </ChakraProvider>
    </BrowserRouter>
  );
}

export default App;
