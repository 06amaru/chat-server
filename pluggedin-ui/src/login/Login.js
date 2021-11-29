import { Container, Box } from '@chakra-ui/layout';
import { Button } from "@chakra-ui/button"
import React, { useState } from 'react'
import { useAuth } from '../auth/UseAuth';
import { Input } from '@chakra-ui/input';
import { useNavigate } from 'react-router';

const Login = () => {
    let context = useAuth()
    const [username, setUsername] = useState()
    const [password, setPassword] = useState()
    const navigate = useNavigate()

    const onLogin = async () => {
      const isAuthenticated = await context.login(username, password)
      if(isAuthenticated) {
        navigate("/", {replace: true})
      }
    }

    return (
        <Container>
            <Box fontSize={30}>Inicia sesión</Box>
            <Box>
                <label>
                  <p>Username</p>
                  <Input type="text" onChange={ (e) => {setUsername(e.target.value)}}/>
                </label>
                <label>
                  <p>Password</p>
                  <Input type="password" onChange={ (e) => {setPassword(e.target.value)}} />
                </label>
                <Box padding={5}>
                  <Button colorScheme="teal" onClick={onLogin}>Iniciar</Button>
                </Box>
            </Box>
        </Container>
    )

}

export default Login;