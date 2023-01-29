import { Accordion, AccordionButton, AccordionIcon, AccordionItem, AccordionPanel, Box, Flex, Text } from "@chakra-ui/react"
import ThirdContent from "../components/landing/thirdcontent";
import HomeNavbar from "../components/navbar/home";
import Sidebar from "../components/sidebar";
import { fontFamilyInter } from "../styling/style";

const HomePage = () => {
    return(
        <Box
        width={{
            'lg' : '100%'
        }}
        marginLeft={{
            'lg' : '16.5%'
        }}
        flexWrap='wrap'>
            <ThirdContent useText={false}/>
            <Box>
                <Box
                marginLeft={{
                    'lg' : '4%'
                }}>
                    <Accordion
                    allowToggle
                    allowMultiple
                    defaultIndex={[0]}>
                        <AccordionItem>
                            <AccordionButton>
                                <Box textAlign='left'>
                                    <Text
                                    fontFamily={fontFamilyInter}>Apa itu My Spend?</Text>
                                </Box>
                                <AccordionIcon
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
                            }}>
                                <Box textAlign='left'
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
                            }}>
                                <Box textAlign='left'
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