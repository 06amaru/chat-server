import React, {useRef, useState} from 'react'
import { Grid, GridItem } from '@chakra-ui/react'
import { Box, VStack, Button } from '@chakra-ui/react'
import { useDisclosure } from '@chakra-ui/hooks'
import Chat from '../chat/Chat'
import Transition from './Transition'

const Home = () => {
    const [chats, setChats] = useState([])
    const { isOpen, onOpen, onClose } = useDisclosure()
    const cancelRef = useRef()
    const [chat, setChat] = useState(null)
    const [receiverId, setReceiverId] = useState("")

    const handleClick = (i) => {
        console.log(i)
        setChat(chats[i])
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
                    <Transition isOpen={isOpen} onClose={onClose} cancelRef={cancelRef} receiverId={receiverId} setReceiverId={setReceiverId} />
                    {
                        chats.length > 0 ? chats.map((c, i) => 
                            <Box key={i} bg="green" width="100%" textAlign="center" onClick={() => handleClick(i)}>
                                {c.receiver}
                            </Box>)
                            : <Box> No tienes ningun chat</Box>
                    }
                </VStack>
            </GridItem>
            <GridItem colSpan={4} bg='tomato' >
                <Box>
                    {
                        chat? <Chat chat={chat} />
                        :<Box>No has seleccionado un chat</Box>
                    }
                    
                </Box>
            </GridItem>
        </Grid>
    )
}

export default Home;