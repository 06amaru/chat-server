import { Button } from "@chakra-ui/button"
import { Box } from "@chakra-ui/layout"

const Messages = (props) => {
    //console.log(props)
    return (
        <Box>
        {
            props.messages.map( (msg, i) => (
            <Button colorScheme="teal" key={i}>
                {msg}
            </Button>
            ))
        }
        </Box>
    )
}

export default Messages