import React, {useEffect, useRef, useState} from 'react'
import { Grid, GridItem } from '@chakra-ui/react'
import { Box, VStack, Button } from '@chakra-ui/react'
import { useDisclosure } from '@chakra-ui/hooks'
import Chat from '../chat/Chat'
import Transition from './Transition'
import eccrypto from "eccrypto"


const Home = () => {
    const [chats, setChats] = useState([])
    const { isOpen, onOpen, onClose } = useDisclosure()
    const cancelRef = useRef()
    const [chat, setChat] = useState(null)
    const [receiverUsername, setReceiverUsername] = useState("")

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
        console.log(i)
        const success = await initChat()
        if (success) {
            setChat(chats[i])
        } else {
            console.log("Hubo problemas")
        }
        
    }

    const initChat = async () => {
        const jwt = localStorage.getItem("jwt")
        const pk = await fetch(`http://127.0.0.1:1323/api/fluent/secret-key-receiver?username=${receiverUsername}`, {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer '+jwt
                }
            })
        const pkJson = await pk.json()
        
        if (pkJson === null) {
            let privateKey = eccrypto.generatePrivate()
            const response = await fetch(`http://127.0.0.1:1323/api/fluent/secret-key-receiver?username=${receiverUsername}`, {
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer '+jwt
                    },
                    body: JSON.stringify({privateKey})
                })
            return response.ok ? true: false
        } else return true
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

    return (
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
                        receiverUsername={receiverUsername} 
                        setReceiverUsername={setReceiverUsername} />
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
                        chat? <Chat chat={chat} receiver={receiverUsername} />
                        :<Box>No has seleccionado un chat</Box>
                    }
                    
                </Box>
            </GridItem>
        </Grid>
    )
}

export default Home;