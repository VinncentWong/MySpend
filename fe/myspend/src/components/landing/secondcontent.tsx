import { Box, Flex, Image, Text } from "@chakra-ui/react";
import { biruMuda, black, blue, fontFamilyInter } from '../../styling/style';

const SecondContent = () => {

    return(
        <Box
        backgroundColor={biruMuda}
        flexWrap='wrap'
        width='100%'
        height='100%'
        >
            <Box
            position='relative'
            width='100%'
            backgroundSize='cover'>
                <Text
                fontFamily={fontFamilyInter}
                color={blue}
                fontWeight={600}
                fontSize={{
                    'sm' : '23px',
                    'lg' : '40px' 
                }}
                width={{
                    'sm' : '80%',
                    'lg' : '40%'
                }}
                position='relative'
                left={{
                    'sm' : '15%',
                    'lg' :' 30%'
                }}
                paddingTop='3%'>Keep your budget update</Text>
            </Box>
            <Flex
            flexDir='row'
            width='100%'
            position='relative'
            gap='3rem'>
                <Flex
                flexDir='column'
                alignItems='center'
                width='30%'
                marginTop={{
                    'sm' : '2rem',
                    'lg' : '7rem'
                }}
                marginLeft='2rem'>
                    <Image
                        src = '/gambar1_landing.jpg'
                        borderRadius='full'
                        objectFit='cover'
                        boxSize={{
                            'lg' : '35%'
                        }}>
                    </Image>
                    <Text
                    fontFamily={fontFamilyInter}
                    color={blue}
                    fontSize={{
                        'lg' : '30px'
                    }}
                    fontWeight={600}
                    marginTop={{
                        'sm' : '2rem',
                        'lg' : '2px'
                    }}>Calculate your budget</Text>
                    <Text
                    fontWeight={400}
                    color={black}
                    fontSize={{
                        'lg' : '14px'
                    }}
                    boxSize={{
                        'lg' : '50%'
                    }}
                    marginTop={{
                        'sm' : '2rem',
                        'lg' : '2px'
                    }}>
                        We bring together all of your accounts, bills and more, 
                        so you can conveniently manage your 
                        finances from one dashboard
                    </Text>
                </Flex>
                <Flex
                flexDir='column'
                alignItems='center'
                width='30%'
                marginTop={{
                    'sm' : '2rem',
                    'lg' : '7rem'
                }}>
                    <Image
                        src = '/gambar2_landing.jpg'
                        borderRadius='full'
                        boxSize={{
                            'lg' : '35%'
                        }}>
                    </Image>
                    <Text
                    fontFamily={fontFamilyInter}
                    color={blue}
                    fontSize={{
                        'lg' : '30px'
                    }}
                    fontWeight={600}
                    marginTop={{
                        'sm' : '2rem',
                        'lg' : '2px'
                    }}>Real-time tracker</Text>
                    <Text
                    fontWeight={400}
                    color={black}
                    fontSize={{
                        'lg' : '14px'
                    }}
                    boxSize={{
                        'lg' : '50%'
                    }}
                    marginTop={{
                        'sm' : '3.5rem',
                        'lg' : '2px'
                    }}>
                       While it’s important that you decide what you want 
                       your money to do now, it’s more important to realize
                    </Text>
                </Flex>
                <Flex
                flexDir='column'
                alignItems='center'
                width='30%'
                marginTop={{
                    'sm' : '2rem',
                    'lg' : '7rem'
                }}
                marginRight='2rem'>
                    <Image
                        src = '/gambar3_landing.jpg'
                        borderRadius='full'
                        objectFit='cover'
                        boxSize={{
                            'lg' : '35%'
                        }}>
                    </Image>
                    <Text
                    fontFamily={fontFamilyInter}
                    color={blue}
                    fontSize={{
                        'lg' : '30px'
                    }}
                    fontWeight={600}
                    marginTop={{
                        'sm' : '2rem',
                        'lg' : '2px'
                    }}>Generate periodic report</Text>
                    <Text
                    fontWeight={400}
                    color={black}
                    fontSize={{
                        'lg' : '14px'
                    }}
                    boxSize={{
                        'lg' : '50%'
                    }}
                    marginTop={{
                        'sm' : '2rem',
                        'lg' : '2px'
                    }}>
                        Make the budgets track the source of problem directly 
                        based on periodic report
                    </Text>
                </Flex>
            </Flex>
        </Box>
    )
}

export default SecondContent;