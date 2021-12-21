import './App.css';
import {ChakraProvider} from '@chakra-ui/react'
import { Route, Navigate, Routes, useLocation, BrowserRouter } from 'react-router-dom';
import { AuthProvider, useAuth } from './auth/UseAuth';
//import { lazy, Suspense } from 'react';
import Login from './components/login/Login';
import Home from './components/home/Home';
import Signup from './signup/Signup';

// const AsyncRoute = ({element, ...props}) => {
//   console.log("async")
//   return <Route {...props} element={element}> </Route>
// }

const AuthenticatedRoute = ({children}) => {
  let { user } = useAuth()
  let location = useLocation() 
  if(user === null) {
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
                  <Home />
                </AuthenticatedRoute>
              } />
              <Route 
                exact
                path="/login"
                element={
                  <Login />
                }
              />
              <Route
                exact
                path="/register"
                element={
                  <Signup />
                }
              />
          </Routes>
        </AuthProvider>
      </ChakraProvider>
    </BrowserRouter>
  );
}

export default App;
