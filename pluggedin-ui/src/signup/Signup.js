import { Button } from '@chakra-ui/button'
import { Input } from '@chakra-ui/input'
import { Box, Container } from '@chakra-ui/layout'
import React, { useState } from 'react'
import eccrypto from "eccrypto"
var CryptoJS = require("crypto-js");


const Signup = () => {
    const [username, setUsername] = useState()
    const [password, setPassword] = useState()

    const signup = async () => {
        const privateKey = eccrypto.generatePrivate()
        const publicKey = eccrypto.getPublic(privateKey)

        const mk = CryptoJS.AES.encrypt( JSON.stringify({
            pk: privateKey
        }), password).toString()
        console.log(publicKey)
        console.log(mk)

        const response = await fetch('http://127.0.0.1:1323/signup', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username,
                password,
                'publickey': publicKey.toString('utf-8'),
                'privatekey': mk
            })
        })

        if (response.status === 200) {
            console.log("register successful")
        } else {
            console.log(response)
        }
    }

    return (
        <Container>
            <Box fontSize={30}>Registrarse</Box>
            <Box>
                <label>
                    <p>Username</p>
                    <Input type="text" onChange={(e) => {setUsername(e.target.value)}} />
                </label>
                <label>
                    <p>Password</p>
                    <Input type="password" onChange={(e) => {setPassword(e.target.value)}} />
                </label>
                <Box>
                    <Button colorScheme="teal" onClick={signup}>DONE</Button>
                </Box>
            </Box>
        </Container>
    )
}

export default Signup;