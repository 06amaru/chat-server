import React, { useEffect } from "react"

import {
    AlertDialog,
    AlertDialogBody,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogContent,
    AlertDialogOverlay,
    Button,
    AlertDialogCloseButton,
    Input
  } from '@chakra-ui/react'

const Transition = (props) => {

    const handleChange = e => {
        props.setReceiverUsername(e.target.value)
    }

    return(
        <>
        <AlertDialog
            motionPreset='slideInBottom'
            leastDestructiveRef={props.cancelRef}
            onClose={props.onClose}
            isOpen={props.isOpen}
            isCentered
        >
            <AlertDialogOverlay />

            <AlertDialogContent>
            <AlertDialogHeader>Por favor ingresar el username de la persona que recibe el mensaje</AlertDialogHeader>
            <AlertDialogCloseButton />
            <AlertDialogBody>
                <Input value={props.receiverUsername} placeholder="username" onChange={handleChange}/>
            </AlertDialogBody>
            <AlertDialogFooter>
                <Button ref={props.cancelRef} onClick={props.onClose}>
                Cancelar
                </Button>
                <Button colorScheme='red' ml={3} onClick={props.createChat}>
                Crear Chat
                </Button>
            </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
        </>
    )
}

export default Transition