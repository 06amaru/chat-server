import { useEffect, useRef, useState } from "react"
import eccrypto from "eccrypto"
import { useAuth } from "../auth/UseAuth"


//other approach to share ws along the app is a context provider but for small app this is ok
const UseWebsocket = (url, chatId, receiver, setMessages) => {
    const ws = useRef(null)
    let context = useAuth()

    useEffect(() => {
        const jwt = localStorage.getItem("jwt")
        if (chatId !== null) {
            ws.current = new WebSocket(url + "?jwt="+jwt+"&id="+chatId)
        } else {
            //create new chat
            //console.log("create new chat")
            ws.current = new WebSocket(url + "?jwt="+jwt+"&receiver="+receiver)
        }
        setMessages([])
        ws.current.onmessage = async (e) => {
            const data = JSON.parse(e.data)
            //console.log(data.sender)
            //console.log(data.body)
            if(data.sender === "Server") {
                setMessages((prev) => [...prev, data.body])
            } else if (data.sender !== context.username) { //solo puede desencriptar el receptor
                //console.log(data.sender)
                //console.log(context.username)
                const body = JSON.parse(data.body)
                //console.log(body.message)
                const encrypted = {
                    ciphertext : Buffer(body.message.ciphertext),
                    ephemPublicKey : Buffer(body.message.ephemPublicKey),
                    iv: Buffer(body.message.iv),
                    mac: Buffer(body.message.mac)
                }
                const user = JSON.parse(context.user)
                //console.log(user)
                //console.log(user.privateKey.data)
                const key = Buffer.from(user.privateKey.data)
                const decrypted = await eccrypto.decrypt(
                    key,
                    encrypted)
                console.log(decrypted.toString())
                setMessages((prev) => [...prev, decrypted.toString()])
            }
            //setMessages((prev) => [...prev, data])
        }

        return () => {
            console.log("cerrando web socket")
            ws.current.close()
            ws.current = null
        }
    }, [url, chatId])

    return {
        ws
    }
} 

export default UseWebsocket;