import {React} from "react"
import { Container } from "@chakra-ui/react"
import InputMessage from "./InputMessage/InputMessage";
import UseWebsocket from "../apis/UseWebsocket";
import Messages from "./Messages/Messages";

const Chat = () => {

    const {ws, messages} = UseWebsocket("ws://localhost:1323/api/plugged/chat\?jwt\=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJvYiIsImV4cCI6MTYzNzQ0OTYyM30.LqXYGFgLPs9lGz6s8X6l8oC2MTJFJQzFfBUbM8UPpzk")
    console.log(messages)
    return (
        <Container>
            <Messages messages={messages} />
            <InputMessage ws={ws}/>
        </Container>
    )
}

export default Chat;