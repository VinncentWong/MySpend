import { Button, Flex, Stack, Text } from "@chakra-ui/react"
import { blue, fontFamilyInter } from '../../styling/style';

const Footer = () => {
    return(
        <Flex
        width='100%'
        height='60vh'
        flexDir='row'
        justifyContent='center'
        backgroundColor={blue}
        backgroundSize='cover'>
            <Stack
            width='50%'
            height='80%'
            marginTop='5%'>
                <Text
                fontFamily={fontFamilyInter}
                fontSize={{
                    'lg' : '46px'
                }}
                color='white'
                fontWeight={600}>Start track your own money</Text>
                <Flex
                flexWrap='wrap'
                flexDir='row'
                height='50%'
                paddingTop='10%'
                justifyContent='center'
                gap={10}>
                    <Button
                    backgroundColor={blue}
                    width='30%'
                    colorScheme='blue'>
                        <Text
                        fontFamily={fontFamilyInter}
                        fontSize={{
                            'lg' : '14px'
                        }}
                        color='white'>See Pricelist</Text>
                    </Button>
                    <Button
                    backgroundColor='#C1D5E8'
                    width='30%'
                    colorScheme='white'>
                        <Text
                        fontFamily={fontFamilyInter}
                        fontSize={{
                            'lg' : '14px'
                        }}
                        color={blue}>Get Started</Text>
                    </Button>
                </Flex>
            </Stack>
        </Flex>
    );
}

export default Footer;