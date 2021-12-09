import {React} from "react"
import { Container } from "@chakra-ui/react"
import InputMessage from "./InputMessage/InputMessage";
import UseWebsocket from "../../apis/UseWebsocket";
import Messages from "./Messages/Messages";

const Chat = (props) => {
    console.log(props)
    const chatId = props.chat !== "nuevo chat" ? props.chat.id : null
    const {ws, messages} = UseWebsocket("ws://localhost:1323/api/plugged/chat", chatId, props.receiver)
    
    return (
        <Container>
            <Messages messages={messages} />
            <InputMessage ws={ws}/>
        </Container>
    )
}

export default Chat;