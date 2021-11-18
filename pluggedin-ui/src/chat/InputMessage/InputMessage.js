import {React, useEffect, useRef} from "react"
import { Input, Box } from "@chakra-ui/react"

const InputMessage = () => {

    const ws = useRef(null)

    useEffect(() => {
        ws.current = new WebSocket("ws://localhost:1323/api/plugged/chat\?jwt\=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imphb2tzIiwiZXhwIjoxNjM3MzAxMTY0fQ.CfWQDNciGotvJ0cKBmvM-KssZN0nmrTDF1FBW7l350A", )
    }, [])

    const onKeyUp = event => {
        if (event.charCode === 13) {
            console.log("SEND MESSAGE")
        }
    }

    return (
        <Box>
            <Input style={{color:"white"}} placeholder="Press enter to send your message" onKeyPress={onKeyUp}/>
        </Box>
    )
}

export default InputMessage;