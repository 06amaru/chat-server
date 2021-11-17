import React from 'react'
import 'react-chat-elements/dist/main.css';
import { MessageList } from 'react-chat-elements'

const History = () => {

    return (
        <div>
            <MessageList 
                dataSource = {[
                    {
                        position: 'right',
                        type: 'text',
                        text: 'hello',
                        date: new Date()
                    },
                    {
                        position: 'right',
                        type: 'text',
                        text: 'world',
                        date: new Date()
                    }
                ]}
            />
        </div>
    )
}

export default History;