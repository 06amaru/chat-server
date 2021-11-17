import {React} from "react"
import { Input, Box } from "@chakra-ui/react"

const InputMessage = () => {

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