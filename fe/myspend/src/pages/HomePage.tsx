import { Accordion, AccordionButton, AccordionIcon, AccordionItem, AccordionPanel, Box, Flex, Text } from "@chakra-ui/react"
import ThirdContent from "../components/landing/thirdcontent";
import HomeNavbar from "../components/navbar/home";
import Sidebar from "../components/sidebar";
import { fontFamilyInter } from "../styling/style";

const HomePage = () => {
    return(
        <Box
        border='4px solid green'
        width={{
            'lg' : '100%'
        }}
        marginLeft={{
            'lg' : '16.5%'
        }}
        flexWrap='wrap'>
            <ThirdContent useText={false}/>
            <Box
            border='4px solid grey'>
                <Box
                border='4px solid red'
                marginLeft={{
                    'lg' : '4%'
                }}>
                    <Accordion
                    allowToggle
                    allowMultiple
                    defaultIndex={[0]}
                    border='4px solid pink'>
                        <AccordionItem>
                            <AccordionButton border='4px solid orange'>
                                <Box textAlign='left'>
                                    <Text
                                    fontFamily={fontFamilyInter}>Apa itu My Spend?</Text>
                                </Box>
                                <AccordionIcon
                                border='4px solid blue'
                                marginLeft={{
                                    'sm' : '65%',
                                    'lg' : '80%'
                                }}/>
                            </AccordionButton>
                            <AccordionPanel pb={4}>
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. 
                                Architecto nulla totam quos tempora delectus, excepturi 
                                repudiandae inventore labore commodi officia consectetur 
                                saepe unde natus assumenda laudantium vitae cupiditate, ad quo!
                            </AccordionPanel>
                        </AccordionItem>
                        <AccordionItem
                        width={{
                            'lg' : '100%'
                        }}>
                            <AccordionButton 
                            width={{
                                'sm' : '100%',
                                'lg' : '100%'
                            }}
                            border='4px solid orange'>
                                <Box textAlign='left'
                                border='4px solid green'
                                marginLeft={{
                                    'sm' : '-1.1%',
                                    'lg' : '-0.5%'
                                }}>
                                    <Text
                                    fontFamily={fontFamilyInter}
                                    width={{
                                        'lg' : '100%'
                                    }}>Apakah My Spend aman digunakan?</Text>
                                </Box>
                                <AccordionIcon
                                border='4px solid blue'
                                marginLeft={{
                                    'sm' : '48%',
                                    'lg' : '65.7%'
                                }}/>
                            </AccordionButton>
                            <AccordionPanel pb={4}>
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. 
                                Architecto nulla totam quos tempora delectus, excepturi 
                                repudiandae inventore labore commodi officia consectetur 
                                saepe unde natus assumenda laudantium vitae cupiditate, ad quo!
                            </AccordionPanel>
                        </AccordionItem>
                        <AccordionItem
                        width={{
                            'lg' : '100%'
                        }}>
                            <AccordionButton 
                            width={{
                                'sm' : '100%',
                                'lg' : '100%'
                            }}
                            border='4px solid orange'>
                                <Box textAlign='left'
                                border='4px solid green'
                                marginLeft={{
                                    'sm' : '-1.1%',
                                    'lg' : '-0.5%'
                                }}>
                                    <Text
                                    fontFamily={fontFamilyInter}
                                    width={{
                                        'lg' : '100%'
                                    }}>Apa saja fitur My Spend?</Text>
                                </Box>
                                <AccordionIcon
                                border='4px solid blue'
                                marginLeft={{
                                    'sm' : '57%',
                                    'lg' : '74.8%'
                                }}/>
                            </AccordionButton>
                            <AccordionPanel pb={4}>
                                Lorem ipsum dolor sit amet consectetur adipisicing elit. 
                                Architecto nulla totam quos tempora delectus, excepturi 
                                repudiandae inventore labore commodi officia consectetur 
                                saepe unde natus assumenda laudantium vitae cupiditate, ad quo!
                            </AccordionPanel>
                        </AccordionItem>
                    </Accordion>
                </Box>
            </Box>
        </Box>
    )
}

export default HomePage;