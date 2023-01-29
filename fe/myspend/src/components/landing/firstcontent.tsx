import { Box, Button, Fade, Image, Text } from "@chakra-ui/react";
import React, { useEffect, useMemo } from "react";
import { fontFamilyInter } from '../../styling/style';
import { ArrowForwardIcon } from "@chakra-ui/icons";
import { useNavigate } from 'react-router-dom';

const FirstContent = () => {

    const textArray = [ 
        'Manage your finance through one website',
        'All your money in one place',
        'Keep your budget update'
    ];
    
    const navigate = useNavigate();
    const [index, setIndex] = React.useState<number>(0);

    useEffect(() => {
        if(index === 3){
            setIndex(0);
            return;
        }
        setTimeout(() => setIndex(index + 1), 2000);  
    }, [index]);

    return(
        <Box 
        background='linear-gradient(360deg, #EDE0F1 60.01%, #EDEFFB 103.33%)'
        height='110vh'
        display='flex'
        flexDir={{
            'sm' : 'column-reverse',
            'lg' : 'row'
        }}
        flexWrap='wrap'>
            <Box
            position='relative'
            width={{
                'sm' : '70%',
                'lg' : '50%'
            }}
            height={{
                'sm' : '25%',
                'lg' : '90%'
            }}
            bottom={{
                'sm' : '7%',
                'lg' : undefined
            }}
            left={{
                'sm' : '20%',
                'lg' : '2%'
            }}
            top={{
                'lg' : '3%'
            }}>
                <Fade in>
                    <Text
                    fontFamily={{fontFamilyInter}}
                    fontWeight={600}
                    color='#215C94'
                    fontSize={{
                        'sm' : '23px',
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
                colorScheme='blue'
                backgroundColor='#215C94'
                position='relative'
                top='80%'
                left='10%'
                width={{
                    'sm' : '60%',
                    'lg' : '30%'
                }}
                onClick={
                    (e) => {
                        navigate("/register");
                    }
                }
                rightIcon={<ArrowForwardIcon color='white'/>}>
                    <Box
                    width='60%'>
                        <Text
                        fontFamily={fontFamilyInter}
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
            blendMode='multiply'
            backgroundRepeat='no-repeat'
            width={{
                'sm' : '100%',
                'lg' : '50%'
            }}
            height={{
                'sm' : '100%',
                'lg' : '100%'
            }}
            backgroundSize={{
                'lg' : 'cover'
            }}
            position='relative'
            right={{
                'sm' : '70%',
                'lg' : '4%'
            }}>
            </Box>
        </Box>
    )
}

export default FirstContent;