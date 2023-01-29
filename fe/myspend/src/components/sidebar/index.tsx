import { Box, Button, Flex, Image, Stack, Text } from "@chakra-ui/react"
import { CrownIcon, DashboardIcon, FolderIcon, HomeIcon } from "../../styling/icon";
import { fontFamilyInter } from "../../styling/style";

const Sidebar = () => {
    return(
        <Stack
        border='4px solid blue'
        spacing={{
            'lg' : '20%'
        }}
        position='fixed'
        opacity={{
            'sm' : '0',
            'lg' : '1'
        }}>
            <Button
            border='4px solid yellow'
            width='100%'>
                <Flex
                border='4px solid grey'
                width='100%'>
                    <Box 
                    border='4px solid green'>
                        <Image
                        src='/home.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    border='4px solid pink'
                    paddingLeft={{
                        'lg' : '20px'
                    }}>Home</Text>
                </Flex>
            </Button>
            <Button
            border='4px solid red'>
                 <Flex
                border='4px solid grey'
                width='100%'>
                    <Box 
                    border='4px solid green'>
                        <Image
                        src='/dashboard.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    border='4px solid pink'
                    paddingLeft={{
                        'lg' : '15px'
                    }}>Dashboard</Text>
                </Flex>
            </Button>
            <Button
            border='4px solid black'>
                 <Flex
                border='4px solid grey'
                width='100%'>
                    <Box 
                    border='4px solid green'>
                        <Image
                        src='/folder.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    border='4px solid pink'
                    paddingLeft={{
                        'lg' : '13px'
                    }}>Manage it</Text>
                    <Box
                    border='4px solid red'>
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
            <Button
            border='4px solid grey'>
                 <Flex
                border='4px solid grey'
                width='100%'>
                    <Box 
                    border='4px solid green'>
                        <Image
                        src='/profile.svg'
                        paddingTop={{
                            'lg' : '5px'
                        }}></Image>
                    </Box> 
                    <Text
                    fontFamily={fontFamilyInter}
                    border='4px solid pink'
                    paddingLeft={{
                        'lg' : '17px'
                    }}>Profile</Text>
                </Flex>
            </Button>
        </Stack>
    )
}

export default Sidebar;