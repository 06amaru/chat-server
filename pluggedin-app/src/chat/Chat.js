import React, { useRef } from 'react';
import { Input, Button } from 'react-chat-elements'
import 'react-chat-elements/dist/main.css';
import History from './history/History';

const Chat =  () => {

    const inputRef = useRef(null)

    const onButtonClick = () => {
        inputRef.current.clear()
    }

    return (
        <div>
            <History />
            <Input 
                ref = {inputRef}
                placeholder = "Write your message..."
                rightButtons = {
                    <Button 
                    color='white'
                    backgroundColor='black'
                    text='Send'
                    onClick={onButtonClick}
                    />
                }
            />
        </div>
    )
}

export default Chat;