import {React} from "react"
import { Container } from "@chakra-ui/react"
import InputMessage from "./InputMessage/InputMessage";
import UseWebsocket from "../../apis/UseWebsocket";
import Messages from "./Messages/Messages";

const Chat = (props) => {
    console.log(props)
    const {ws, messages} = UseWebsocket("ws://localhost:1323/api/plugged/chat")
    
    return (
        <Container>
            <Messages messages={messages} />
            <InputMessage ws={ws}/>
        </Container>
    )
}

export default Chat;