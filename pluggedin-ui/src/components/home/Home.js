import React, {useState} from 'react'
import { Grid, GridItem } from '@chakra-ui/react'

const Home = () => {

    return (
        <Grid
        h="100%"
        templateRows='repeat(1, 1fr)'
        templateColumns='repeat(5, 1fr)'
        gap={4}
        >
            <GridItem rowSpan={1} colSpan={1} bg='red' />
            <GridItem colSpan={4} bg='tomato' />
        </Grid>
    )
}

export default Home;