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
        props.setReceiverId(e.target.value)
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
            <AlertDialogHeader>Por favor ingresar ID del receptor</AlertDialogHeader>
            <AlertDialogCloseButton />
            <AlertDialogBody>
                <Input value={props.receiverId} placeholder="ID de receptor. eg: 666" onChange={handleChange}/>
            </AlertDialogBody>
            <AlertDialogFooter>
                <Button ref={props.cancelRef} onClick={props.onClose}>
                Cancelar
                </Button>
                <Button colorScheme='red' ml={3}>
                Crear
                </Button>
            </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
        </>
    )
}

export default Transition