import { Box, Button, Fade, Image, Text } from "@chakra-ui/react";
import React, { useEffect, useMemo } from "react";
import { fontFamily } from '../../styling/style';
import { ArrowForwardIcon } from "@chakra-ui/icons";

const FirstContent = () => {

    const font = 'Inter, sans-serif';

    const textArray = [ 
        'Manage your finance through one website',
        'All your money in one place',
        'Keep your budget update'
    ];

    const [index, setIndex] = React.useState<number>(0);

    useEffect(() => {
        if(index === 3){
            setIndex(0);
        }
        setTimeout(() => setIndex(index + 1), 2000);  
    }, [index]);

    return(
        <Box 
        background='linear-gradient(360deg, #EDE0F1 60.01%, #EDEFFB 103.33%)'
        height='100%'
        display='flex'
        flexDir='row'
        flexWrap='wrap'>
            <Box
            position='relative'
            width='50%'>
                <Fade in>
                    <Text
                    fontFamily={{font}}
                    fontWeight={600}
                    color='#215C94'
                    fontSize={{
                        'lg' : '55px'
                    }}
                    position='absolute'
                    top='30%'
                    left='9%'
                    width='70%'>
                        {textArray[index]}
                    </Text>
                </Fade>
                <Button
                backgroundColor='#215C94'
                position='relative'
                top='80%'
                left='10%'
                width='30%'
                rightIcon={<ArrowForwardIcon color='white'/>}>
                    <Box
                    width='60%'>
                        <Text
                        fontFamily={fontFamily}
                        fontSize={{
                            'lg' : '14px'
                        }}
                        fontWeight={600}
                        color='white'>Get started</Text>
                    </Box>
                </Button>
            </Box>
            <Box
            backgroundImage='./hp.jpg'
            zIndex='2'
            blendMode='multiply'
            backgroundRepeat='no-repeat'
            boxSize='45%'
            backgroundSize='cover'
            position='relative'
            right='4%'>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <br></br>
            </Box>
        </Box>
    )
}

export default FirstContent;