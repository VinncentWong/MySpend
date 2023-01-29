import { Box, Button, Flex, Image, Stack, Text } from "@chakra-ui/react";
import { blue, fontFamilyInter } from "../../styling/style";

const HomeNavbar = () => {
    return(
        <Box
        width={{
            'lg' : '100%'
        }}
        height={{
            'lg' : '100%'
        }}>
            <Flex>
                <Flex
                flexDir='row'
                flexWrap='wrap'
                width={{
                    'lg' : '21%'
                }}
                marginLeft={{
                    'lg' : '2rem'
                }}>
                    <Image
                    src="/Logo.svg"/>
                    <Text
                    fontFamily={fontFamilyInter}
                    color={blue}
                    fontSize={{
                        'lg' : '35px'
                    }}
                    fontWeight='600'
                    paddingTop={{
                        'sm' : '10%',
                        'lg' : '0'
                    }}>My Spend</Text>
                </Flex>
                <Box
                width={{
                    'lg' : '10%'
                }}
                marginLeft={{
                    'sm' : '1rem',
                    'lg' : '2rem'
                }}>
                    <Text
                    fontFamily={fontFamilyInter}
                    fontSize={{
                        'lg' : '35px'
                    }}
                    fontWeight='600'
                    paddingTop={{
                        'sm' : '40%',
                        'lg' : '7%'
                    }}>Home</Text>
                </Box>
                <Flex
                width={{
                    'sm' : '40%',
                    'lg' : '15%'
                }}
                flexDir='row'
                flexWrap='wrap'
                marginLeft={{
                    'sm' : '5%',
                    'lg' : '50%'
                }}
                gap={{
                    'sm' : '15%',
                    'lg' : '30%'
                }}
                alignItems='center'>
                    <Button>
                        <Image
                        src='settings.svg'
                        marginLeft={{
                            'lg' : '10%'
                        }}></Image>
                    </Button>
                    <Button>
                        <Image
                        src='bell.svg'></Image>
                    </Button>
                </Flex>
            </Flex>
        </Box>
    )
}

export default HomeNavbar;