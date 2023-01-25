import { Box, Flex, Image, Text } from "@chakra-ui/react";

const SecondContent = () => {

    const blue = '#215C94';
    const black = '#1A2127';
    const biruMuda = '#EDEFFB';
    const fontFamily = 'Inter, sans-serif';

    return(
        <Box
        position='absolute'
        backgroundColor={biruMuda}
        flexWrap='wrap'
        >
            <Text
            fontFamily={fontFamily}
            color={blue}
            fontWeight={600}
            fontSize={{
                'lg' : '40px' 
            }}
            marginTop='2rem'
            width='50%'
            marginLeft='27%'>Keep your budget update</Text>
            <Flex
            flexDir='row'
            width='100%'
            position='relative'
            gap='3rem'>
                <Flex
                flexDir='column'
                alignItems='center'
                width='30%'
                marginTop='13rem'
                marginLeft='2rem'>
                    <Image
                        src = '/gambar1_landing.jpg'
                        borderRadius='full'
                        objectFit='cover'
                        boxSize='35%'>
                    </Image>
                    <Text
                    fontFamily={fontFamily}
                    color={blue}
                    fontSize={{
                        'lg' : '30px'
                    }}
                    fontWeight={600}
                    marginTop={'2px'}>Calculate your budget</Text>
                    <Text
                    fontWeight={400}
                    color={black}
                    fontSize={{
                        'lg' : '14px'
                    }}
                    boxSize={{
                        'lg' : '50%'
                    }}
                    marginTop={'2px'}>
                        We bring togeher all of your accounts, bills and more, 
                        so you can conveniently manage your 
                        finances from one dashboard
                    </Text>
                </Flex>
                <Flex
                flexDir='column'
                alignItems='center'
                width='30%'
                marginTop='13rem'>
                    <Image
                        src = '/gambar2_landing.jpg'
                        borderRadius='full'
                        boxSize='35%'>
                    </Image>
                    <Text
                    fontFamily={fontFamily}
                    color={blue}
                    fontSize={{
                        'lg' : '30px'
                    }}
                    fontWeight={600}
                    marginTop={'2px'}>Real-time tracker</Text>
                    <Text
                    fontWeight={400}
                    color={black}
                    fontSize={{
                        'lg' : '14px'
                    }}
                    boxSize={{
                        'lg' : '50%'
                    }}
                    marginTop={'2px'}>
                       While it’s important that you decide what you want 
                       your money to do now, it’s more important to realize
                    </Text>
                </Flex>
                <Flex
                flexDir='column'
                alignItems='center'
                width='30%'
                marginTop='13rem'
                marginRight='2rem'>
                    <Image
                        src = '/gambar3_landing.jpg'
                        borderRadius='full'
                        objectFit='cover'
                        boxSize='35%'>
                    </Image>
                    <Text
                    fontFamily={fontFamily}
                    color={blue}
                    fontSize={{
                        'lg' : '30px'
                    }}
                    fontWeight={600}
                    marginTop={'2px'}>Generate periodic report</Text>
                    <Text
                    fontWeight={400}
                    color={black}
                    fontSize={{
                        'lg' : '14px'
                    }}
                    boxSize={{
                        'lg' : '50%'
                    }}
                    marginTop={'2px'}>
                        Make the budgets track the source of problem directly 
                        based on periodic report
                    </Text>
                </Flex>
            </Flex>
        </Box>
    )
}

export default SecondContent;