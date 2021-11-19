import {React} from "react"
import { Container } from "@chakra-ui/react"
import InputMessage from "./InputMessage/InputMessage";
import UseWebsocket from "../apis/UseWebsocket";
import Messages from "./Messages/Messages";

const Chat = () => {

    const {ws, messages} = UseWebsocket("ws://localhost:1323/api/plugged/chat\?jwt\=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imphb2tzIiwiZXhwIjoxNjM3Mzg3NjY5fQ.Dt6HWJ6WwqI-V8P_mf9F6-2VzM10zH7g9mLYcJvH33I")
    console.log(messages)
    return (
        <Container>
            <Messages messages={messages} />
            <InputMessage ws={ws}/>
        </Container>
    )
}

export default Chat;