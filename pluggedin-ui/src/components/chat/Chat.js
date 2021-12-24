import {React, useState} from "react"
import { Container } from "@chakra-ui/react"
import InputMessage from "./InputMessage/InputMessage";
import UseWebsocket from "../../apis/UseWebsocket";
import Messages from "./Messages/Messages";

const Chat = (props) => {
    //console.log(props)
    const chatId = props.chat !== "nuevo chat" ? props.chat.id : null
    //console.log(chatId)
    const [messages, setMessages] = useState([])
    const {ws} = UseWebsocket("ws://localhost:1323/api/plugged/chat", chatId, props.receiver, setMessages)
    
    return (
        <Container>
            <Messages messages={messages} />
            <InputMessage ws={ws} publicKey={props.publicKey} setMessages={setMessages}/>
        </Container>
    )
}

export default Chat;