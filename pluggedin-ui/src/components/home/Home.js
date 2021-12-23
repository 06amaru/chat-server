import React, {useEffect, useRef, useState} from 'react'
import { Grid, GridItem } from '@chakra-ui/react'
import { Box, VStack, Button } from '@chakra-ui/react'
import { Flex, Spacer } from '@chakra-ui/layout'
import { useDisclosure } from '@chakra-ui/hooks'
import Chat from '../chat/Chat'
import Transition from './Transition'
import { useAuth } from '../../auth/UseAuth'
import { useNavigate } from 'react-router';
import { base64 } from "rfc4648";

const Home = () => {
    const [chats, setChats] = useState([])
    const { isOpen, onOpen, onClose } = useDisclosure()
    const cancelRef = useRef()
    const [chat, setChat] = useState(null)
    const [username, setusername] = useState("")
    const [publicKey, setPublicKey] = useState(null)
    let context = useAuth()
    const navigate = useNavigate()

    useEffect(() => {
        async function getChats() {
            const jwt = localStorage.getItem("jwt")
            const response = await fetch('http://127.0.0.1:1323/api/fluent/chats', {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer '+jwt
                }
            })
            if (response.status === 200) {
                let data = await response.json()
                console.log(...data)
                setChats(data)
            }
        }

        getChats()
    }, [])

    const handleClick = async (i) => {
        const success = await fetchChat(chats[i])
        if (success) {
            setChat(chats[i])
        } else {
            console.log("Hubo problemas")
        }
        
    }

    const fetchChat = async (chat) => {
        //get usernames from chat
        //then get public key
        const jwt = localStorage.getItem("jwt")
        console.log(chat.id)
        const response = await fetch(`http://127.0.0.1:1323/api/fluent/members?chatID=${chat.id}`, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer '+jwt
            }
        })
        const members = await response.json()
        console.log(members)
        try {
            return false
        } catch (error) {
            
        }
    }

    const initChat = async () => {
        try {
            const jwt = localStorage.getItem("jwt")
            const pk = await fetch(`http://127.0.0.1:1323/api/fluent/public-key?username=${username}`, {
                    method: 'GET',
                    headers: {
                        'Authorization': 'Bearer '+jwt
                    }
                })
            const pkJson = await pk.json()
            setPublicKey(base64.parse(pkJson))
            return true    
        } catch (error) {
            console.log(error)
            return false
        }
        
    }

    const createChat = async () => {
        const success = await initChat()
        if (success) {
            setChat("nuevo chat")
            onClose()
        } else {
            console.log("Hubo problemas")
        }
    }

    const onLogout = () => {
        context.logout()
        navigate("/login", {replace: true})
    }

    return (
        <>
        <Flex>
            <Box>Bienvenido</Box>
            <Spacer />
            <Button onClick={onLogout}>Cerrar Sesión</Button>
        </Flex>
        <Grid
        h="100%"
        templateRows='repeat(1, 1fr)'
        templateColumns='repeat(5, 1fr)'
        gap={4}
        >
            <GridItem rowSpan={1} colSpan={1} bg='red'>
                <VStack spacing={4}>
                    <Button onClick={onOpen}>Crear chat</Button>
                    <Transition isOpen={isOpen} 
                        createChat={createChat}
                        onClose={onClose}
                        cancelRef={cancelRef} 
                        username={username} 
                        setusername={setusername} />
                    {
                        chats.length > 0 ? chats.map((c, i) => 
                            <Box key={i} bg="green" width="100%" textAlign="center" onClick={() => handleClick(i)}>
                                {c.id}
                            </Box>)
                            : <Box> No tienes ningun chat</Box>
                    }
                </VStack>
            </GridItem>
            <GridItem colSpan={4} bg='tomato' >
                <Box>
                    {
                        chat? <Chat chat={chat} receiver={username} publicKey={publicKey}/>
                        :<Box>No has seleccionado un chat</Box>
                    }
                    
                </Box>
            </GridItem>
        </Grid>
        </>
    )
}

export default Home;