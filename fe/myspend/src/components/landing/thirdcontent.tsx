import { ArrowForwardIcon } from "@chakra-ui/icons";
import { Box, Button, Card, CardBody, CardFooter, CardHeader, Flex, Heading, Image, Stack, Text } from "@chakra-ui/react"
import { blue, fontFamilyInter } from '../../styling/style';

const ThirdContent = () => {

    return(
        <Box>
            <Box
            width='53%'
            marginLeft='26%'
            marginTop='3%'>
                <Text
                color={blue}
                fontFamily={fontFamilyInter}
                fontWeight={600}
                fontSize={{
                    'lg' : '46px'
                }}>
                    Take a deeper dive into a 
                    new way to track your money</Text>
            </Box>
            <Flex
            flexDir='row'
            width='100%'
            height='100%'
            gap={10}
            marginTop='7%'>
                <Flex
                flexDir='column'
                flexWrap='wrap'
                width='30%'
                marginLeft='5%'>
                    <Card
                    direction={{
                        'lg' : 'column'
                    }}
                    width='90%'
                    height='95%'
                    backgroundColor={blue}
                    marginLeft='5%'
                    marginTop='3%'
                    >
                        <CardHeader
                        fontFamily={fontFamilyInter}
                        color='white'
                        fontSize={{
                            'lg' : '16px'
                        }}
                        marginLeft='9%'>Collection</CardHeader>
                        <CardBody>
                            <Heading 
                            fontFamily={fontFamilyInter}
                            color='white'
                            width='80%'
                            fontSize='24px'
                            marginLeft='10%'
                            marginBottom='3%'
                            fontWeight='700'>
                                My Spend as your personal finance management
                            </Heading>
                            <Flex
                            width='70%'
                            borderRadius='16px'
                            backgroundColor='rgba(193, 213, 232, 0.3);'
                            flexDir='column'
                            alignItems='center'
                            justifyContent='flex-start'
                            marginLeft='10%'
                            >
                                <Image
                                    src='content2_1.png'
                                    width='100%'>
                                </Image>
                            </Flex>
                        </CardBody>
                        <CardFooter
                        display='flex'
                        flexDir='row'
                        justifyContent='flex-end'>
                            <Button 
                            rightIcon={<ArrowForwardIcon color='white'/>}
                            backgroundColor={blue}
                            colorScheme='blue'>
                                <Text
                                fontFamily={fontFamilyInter}
                                fontSize={{
                                    'lg' : '14px'
                                }}
                                color='white'>See All</Text>
                            </Button>
                        </CardFooter>
                    </Card>
                </Flex>
                <Flex
                flexDir='column'
                flexWrap='wrap'
                width='25%'
                alignItems='center'
                gap={3}>
                    <Image
                    src='/chart.png'
                    width='80%'></Image>
                    <Stack gap={6}
                    alignItems='center'>
                        <Text
                        fontSize={{
                            'lg' : '14px'
                        }}
                        fontFamily={fontFamilyInter}
                        fontWeight='500'
                        width='50%'
                        alignSelf='flex-start'>ARTICLE</Text>
                        <Text
                        fontSize={{
                            'lg' : '30px'
                        }}
                        fontFamily={fontFamilyInter}
                        fontWeight='600'
                        textOverflow='ellipsis'
                        width='100%'
                        alignSelf='flex-start'>
                            Cara Mengatur Gaji Pakai Metode 50/30/20, Agar Cashflow Terarah
                        </Text>
                        <Button 
                        rightIcon={<ArrowForwardIcon color={blue}/>} 
                        alignSelf='flex-end'
                        colorScheme='white'
                        >
                            <Text
                            fontFamily={fontFamilyInter}
                            fontSize={{
                                'lg' : '14px'
                            }}
                            color={blue}>Read more</Text>
                        </Button>
                    </Stack>
                </Flex>
                <Flex
                flexDir='column'
                flexWrap='wrap'
                width='25%'
                alignItems='center'
                gap={3}>
                    <Image
                    src='/chart_bro.png'
                    width='80%'></Image>
                    <Stack gap={6}
                    alignItems='center'
                    width='100%'
                    height='60%'>
                        <Text
                        fontSize={{
                            'lg' : '14px'
                        }}
                        fontFamily={fontFamilyInter}
                        fontWeight='500'
                        width='50%'
                        alignSelf='flex-start'>ARTICLE</Text>
                        <Text
                        fontSize={{
                            'lg' : '30px'
                        }}
                        fontFamily={fontFamilyInter}
                        fontWeight='600'
                        textOverflow='ellipsis'
                        width='100%'
                        alignSelf='flex-start'>
                            Yuk Lebih Melek Finansial dengan 4 Cara Ini!
                        </Text>
                        <Box
                        position='relative'
                        top='14%'
                        alignSelf='flex-end'>
                            <Button 
                            rightIcon={<ArrowForwardIcon color={blue}/>} 
                            alignSelf='flex-end'
                            colorScheme='white'>
                                <Text
                                fontFamily={fontFamilyInter}
                                fontSize={{
                                    'lg' : '14px'
                                }}
                                color={blue}>Read more</Text>
                            </Button>
                        </Box>
                    </Stack>
                </Flex>
            </Flex>
        </Box>
    )
}

export default ThirdContent;