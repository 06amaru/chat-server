import { useEffect, useRef, useState } from "react"
import eccrypto from "eccrypto"


//other approach to share ws along the app is a context provider but for small app this is ok
const UseWebsocket = (url, chatId, receiver) => {
    const [messages, setMessages] = useState([])
    const ws = useRef(null)

    useEffect(() => {
        const jwt = localStorage.getItem("jwt")
        if (receiver === "") {
            ws.current = new WebSocket(url + "?jwt="+jwt+"&id="+chatId)
        } else {
            //create new chat
            console.log("create new chat")
            ws.current = new WebSocket(url + "?jwt="+jwt+"&receiver="+receiver)
        }
        ws.current.onmessage = async (e) => {
            const data = JSON.parse(e.data)
            console.log(data.sender)
            console.log(data.body)
            if(data.sender === "Server") {
                setMessages((prev) => [...prev, data.body])
            } else {
                const body = JSON.parse(data.body)
                //console.log(body.pk)
                //console.log(body.message)
                const publicKey = {
                    ciphertext : Buffer(body.message.ciphertext),
                    ephemPublicKey : Buffer(body.message.ephemPublicKey),
                    iv: Buffer(body.message.iv),
                    mac: Buffer(body.message.mac)
                }
                //console.log(publicKey)
                const decrypted = await eccrypto.decrypt(Buffer(body.pk), publicKey)
                //console.log(decrypted)
                setMessages((prev) => [...prev, decrypted.toString()])
            }
            //setMessages((prev) => [...prev, data])
        }

        return () => {
            console.log("cerrando web socket")
            ws.current.close()
            ws.current = null
        }
    }, [url])

    return {
        ws,
        messages
    }
} 

export default UseWebsocket;