import {React} from "react"
import { Container } from "@chakra-ui/react"
import InputMessage from "./InputMessage/InputMessage";
import UseWebsocket from "../apis/UseWebsocket";
import Messages from "./Messages/Messages";

const Chat = () => {

    const {ws, messages} = UseWebsocket("ws://localhost:1323/api/plugged/chat")
    console.log(messages)
    return (
        <Container>
            <Messages messages={messages} />
            <InputMessage ws={ws}/>
        </Container>
    )
}

export default Chat;