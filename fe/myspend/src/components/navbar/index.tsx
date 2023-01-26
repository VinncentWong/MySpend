import { Box, Flex, Image, Link, Text } from "@chakra-ui/react"
import { fontFamilyNunito as fontFamily } from "../../styling/style";
import { textSize } from "../../styling/textsize";
const black = '#1A2127';

const Navbar = () => {
    return(
        <Flex 
        flexDir='row'
        backgroundColor='#EDE0F1'
        flexWrap='wrap'>
            <Box 
            width='15%'
            height='100%'
            position='relative'
            left='10%'>
                <Link 
                display='flex' 
                flexDir='row'
                alignItems='center'>
                    <Image src='Logo.svg'></Image>
                    <Text
                    fontFamily={fontFamily}
                    alignSelf='center'
                    fontSize={textSize}
                    color='#215C94'
                    fontWeight='bold'>MySpend</Text>
                </Link>
            </Box>
            <Box
            position='relative'
            left='30%'
            display='flex'
            flexDir='row'
            gap={{
                'lg' : '4rem'
            }}
            alignItems={'center'}
            width='30%'
            flexWrap='wrap'>
                <Link><Text
                fontFamily={fontFamily}
                fontSize={textSize}
                color={black}>Features</Text></Link>

                <Link><Text
                fontFamily={fontFamily}
                fontSize={textSize}
                color={black}>FAQ</Text></Link>

                <Link><Text
                fontFamily={fontFamily}
                fontSize={textSize}
                color={black}>About</Text></Link>
            </Box>
        </Flex>
    )
}

export default Navbar;