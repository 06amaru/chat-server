import {React, useState} from "react"
import { Input, Box } from "@chakra-ui/react"
import eccrypto from "eccrypto"
//import { useAuth } from "../../../auth/UseAuth";
var CryptoJS = require("crypto-js");

const InputMessage = (props) => {
    //let context = useAuth()
    const [message, setMessage] = useState("")

    const onKeyUp = async(event) => {
        if (event.charCode === 13 && message !== "") {
            const key = 'random-key'
            props.setMessages((prev) => [...prev, message])
            let msgEncrypted = CryptoJS.AES.encrypt(message, key).toString()
            const encrypted = await eccrypto.encrypt(
                Buffer.from(props.publicKey), 
                Buffer.from(msgEncrypted+":::::"+key))
            //console.log(encrypted)
            props.ws.current.send(
                JSON.stringify({
                    "message": encrypted
                })
            )
            //const user = JSON.parse(context.user)
            //const privateKey = Buffer.from(user.privateKey.data)
            //const decrypted = await eccrypto.decrypt(privateKey, encrypted)
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
                value={message}
                placeholder="Press enter to send your message" 
                onKeyPress={onKeyUp}
                onChange={handleChange}
                />
        </Box>
    )
}

export default InputMessage;