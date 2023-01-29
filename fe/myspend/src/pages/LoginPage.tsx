import { Box, Button, Card, CardBody, CardFooter, CardHeader, Checkbox, Flex, FormControl, FormErrorMessage, FormHelperText, FormLabel, Image, Input, InputGroup, InputRightElement, Link, Text } from "@chakra-ui/react"
import { ChangeEvent, FormEvent, MouseEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api";
import { blue, fontFamilyInter, fontFamilyNunito } from '../styling/style';

const LoginPage = () => {
    const [password, setPassword] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [show, setShow] = useState<boolean>(false);
    const navigate = useNavigate();

    const passwordHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value);
    }

    const emailHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    }

    const showHandler = (e: MouseEvent<HTMLButtonElement>) => {
        setShow(!show);
    }


    const submitHandler = async (e: FormEvent<HTMLButtonElement>) => {
       const response = await api.post("/user/login", {
        "email" : email,
        "password" : password
       })
       localStorage.setItem("user", JSON.stringify(response));
       navigate("/home");
    }

    return(
        <Flex
        width='100%'
        height='100%'
        backgroundImage='/background_auth.svg'
        backgroundSize='cover'
        position='relative'
        flexDir={{
            'sm' : 'column',
            'lg' : 'row'
        }}
        flexWrap='wrap'>
            <Flex
            position='absolute'
            flexWrap='wrap'
            flexDir='row'
            width={{
                'sm' : '50%',
                'lg' : '25%'
            }}
            marginLeft={{
                'lg' : '7%'
            }}
            alignSelf={{
                'sm' : 'center',
                'lg' : 'flex-end'
            }}
            marginBottom={{
                'lg' : '20%'
            }}
            top={{
                'lg' : '55%'
            }}
            marginTop={{
                'sm' : '5%',
                'lg' : '0'
            }}>
                <Box
                backgroundColor='white'
                borderRadius='14px'
                height='15%'
                marginTop='5%'>
                    <Image src='/Logo.svg'></Image>
                </Box>
                <Text
                fontFamily={fontFamilyNunito}
                fontWeight='700'
                paddingTop={{
                    'sm' : '10%',
                    'lg' : '6%'
                }}
                fontSize={{
                    'sm' : '25px',
                    'lg' : '40px'
                }}
                color='white'
                marginLeft={{
                    'sm' : '5%'
                }}>My Spend</Text>
            </Flex>
            <Card
            width={{
                'lg' : '50%'
            }}
            borderRadius='14px'
            position='relative'
            marginTop={{
                'sm' : '30%',
                'lg' : '5%'
            }}
            left={{
                'lg' : '40%'
            }}
            height={{
                'lg' : '85%'
            }}
            marginBottom={{
                'sm' : '70%',
                'lg' : '15%'
            }}>
                <CardHeader>
                    <Text
                    fontFamily={fontFamilyInter}
                    fontSize={{
                        'sm' : '23px',
                        'lg' : '34px'
                    }}
                    fontWeight='600'>Welcome back!</Text>
                </CardHeader>
                <CardBody
                alignItems='center'>
                    <Flex
                    flexDir={{
                        'lg' : 'column'
                    }}
                    flexWrap='wrap'
                    gap='1rem'>
                         <Box width='100%'>
                            <FormControl>
                                <FormLabel
                                fontFamily={fontFamilyInter}
                                fontSize={{
                                    'lg' : '16px'
                                }}
                                fontWeight='500'>Email</FormLabel>
                                <Input 
                                variant='filled'
                                placeholder='Enter your email'
                                value={email}
                                onChange={emailHandler}
                                type='email'></Input>
                            </FormControl>
                        </Box>
                        <Box width='100%'>
                            <FormControl>
                                <FormLabel
                                fontFamily={fontFamilyInter}
                                fontSize={{
                                    'lg' : '16px'
                                }}
                                fontWeight='500'>Password</FormLabel>
                                <InputGroup>
                                    <Input 
                                    variant='filled'
                                    value={password}
                                    onChange={passwordHandler}
                                    placeholder='Enter your password'
                                    type={show? 'text' : 'password'}></Input>
                                    <InputRightElement
                                    width={{
                                        'lg' : '15%'
                                    }}>
                                        <Button
                                        onClick={showHandler}>
                                            { 
                                                show ? 
                                                <Text
                                                fontFamily={fontFamilyInter}
                                                fontSize={{
                                                    'lg' : '14px'
                                                }}>Hide</Text> 
                                                : 
                                                <Text
                                                fontFamily={fontFamilyInter}
                                                fontSize={{
                                                    'lg' : '14px'
                                                }}>Show</Text>
                                            }
                                        </Button>
                                    </InputRightElement>
                                </InputGroup>
                            </FormControl>
                        </Box>
                    </Flex>
                </CardBody>
                <CardFooter
                flexDir='row'>
                    <Box>
                        <Text
                        fontFamily={fontFamilyInter}>New user? <Link 
                            color={blue} 
                            fontFamily={fontFamilyInter}
                            href="/register">
                                Sign up
                            </Link>
                        </Text>
                    </Box>
                    <Button
                    backgroundColor={blue}
                    colorScheme='blue'
                    marginLeft={{
                        'sm' : '25%',
                        'lg' : '45%'
                    }}
                    onClick={submitHandler}>
                        <Text
                        fontFamily={fontFamilyInter}
                        fontSize={{
                            'lg' : '14px'
                        }}
                        fontWeight='600'>Login</Text>
                    </Button>
                </CardFooter>
            </Card>
        </Flex>
    )
}

export default LoginPage;