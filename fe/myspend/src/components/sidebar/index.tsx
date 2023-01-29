import { Box, Button, Flex, Image, Stack, Text } from "@chakra-ui/react"
import { CrownIcon, DashboardIcon, FolderIcon, HomeIcon } from "../../styling/icon";
import { fontFamilyInter } from "../../styling/style";

const Sidebar = () => {
    return(
        <Stack
        spacing={{
            'lg' : '20%'
        }}
        position='fixed'
        opacity={{
            'sm' : '0',
            'lg' : '1'
        }}>
            <Button
            width='100%'>
                <Flex
                width='100%'>
                    <Box>
                        <Image
                        src='/home.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    paddingLeft={{
                        'lg' : '20px'
                    }}>Home</Text>
                </Flex>
            </Button>
            <Button>
                 <Flex
                width='100%'>
                    <Box>
                        <Image
                        src='/dashboard.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    paddingLeft={{
                        'lg' : '15px'
                    }}>Dashboard</Text>
                </Flex>
            </Button>
            <Button>
                 <Flex
                width='100%'>
                    <Box>
                        <Image
                        src='/folder.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    paddingLeft={{
                        'lg' : '13px'
                    }}>Manage it</Text>
                    <Box>
                        <Image
                        src='/crown.svg'
                        paddingLeft={{
                            'lg' : '10px'
                        }}
                        paddingTop={{
                            'lg' : '3px'
                        }}/>
                    </Box>
                </Flex>
            </Button>
            <Button>
                 <Flex
                width='100%'>
                    <Box>
                        <Image
                        src='/profile.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    paddingLeft={{
                        'lg' : '17px'
                    }}>Profile</Text>
                </Flex>
            </Button>
        </Stack>
    )
}

export default Sidebar;