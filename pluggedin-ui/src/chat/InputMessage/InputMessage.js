import {React, useEffect, useRef, useState} from "react"
import { Input, Box } from "@chakra-ui/react"
import eccrypto from "eccrypto"

const InputMessage = (props) => {

    const [message, setMessage] = useState("")

    const onKeyUp = async(event) => {
        if (event.charCode === 13 && message !== "") {
            const privateKey = eccrypto.generatePrivate()
            const publicKey = eccrypto.getPublic(privateKey)
            const encrypted = await eccrypto.encrypt(publicKey, Buffer.from(message))

            props.ws.current.send(
                JSON.stringify({
                    "message": encrypted,
                    "pk": privateKey
                })
            )
            //const decrypted = await eccrypto.decrypt(privateKey, encrypted)
            //console.log("A?")
            //console.log(decrypted.toString())
            setMessage("")
        }
    }

    const handleChange = event => {
        setMessage(event.target.value)
    }

    return (
        <Box>
            <Input 
                style={{color:"white"}} 
                value={message}
                placeholder="Press enter to send your message" 
                onKeyPress={onKeyUp}
                onChange={handleChange}
                />
        </Box>
    )
}

export default InputMessage;