import { Box, Button, Card, CardBody, CardFooter, CardHeader, Checkbox, Flex, FormControl, FormErrorMessage, FormHelperText, FormLabel, Image, Input, InputGroup, InputRightElement, Link, Text } from "@chakra-ui/react"
import { string } from "prop-types";
import { ChangeEvent, FormEvent, MouseEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api";
import { blue, fontFamilyInter, fontFamilyNunito } from '../styling/style';

const RegisterPage = () => {

    const expression: RegExp = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;

    const [name, setName] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [show, setShow] = useState<boolean>(false);
    const [notcheck, setNotCheck] = useState<boolean>(true);
    const navigate = useNavigate();

    const isNameValid = name.length >= 6;
    const isPasswordValid = password.length >= 6;
    const isEmailValid = expression.test(email) && email.length >= 1;

    const nameHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setName(e.target.value);
    }

    const passwordHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value);
    }

    const emailHandler = (e: ChangeEvent<HTMLInputElement>) => {
        setEmail(e.target.value);
    }

    const showHandler = (e: MouseEvent<HTMLButtonElement>) => {
        setShow(!show);
    }

    const checkHandler = (e: ChangeEvent<HTMLInputElement>) => {
        console.log("check terpanggil dengan nilai ", notcheck)
        setNotCheck(!notcheck);
    }

    const submitHandler = (e: FormEvent<HTMLButtonElement>) => {
        console.log(`terpanggil dengan
        isNameValid = ${isNameValid}
        isPasswordValid = ${isPasswordValid}
        isEmailValid = ${isEmailValid}
        notcheck = ${!notcheck}
        `)
        if(isNameValid && isPasswordValid && isEmailValid && !notcheck){
            // call API
            const response =  api.post("/user/register", {
                "name" : name,
                "email" : email,
                "password" : password
            });
            response.then((data) => {
                console.log(JSON.stringify(data));
                navigate("/");
            }).catch((err) => {
                console.log(JSON.stringify(err));
            });
            e.preventDefault();
        } else {
            e.preventDefault();
        }
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
            flexWrap='wrap'
            flexDir='row'
            width={{
                'sm' : '50%',
                'lg' : '23%'
            }}
            marginLeft={{
                'lg' : '7%'
            }}
            height={{
                'lg' : '100%'
            }}
            alignSelf={{
                'sm' : 'center',
                'lg' : 'flex-end'
            }}
            marginTop={{
                'sm' : '3%',
                'lg' : '0'
            }}
            marginBottom={{
                'sm' : '10%',
                'lg' : '20%'
            }}>
                <Box
                backgroundColor='white'
                borderRadius='14px'
                height={{
                    'lg' : '20%'
                }}>
                    <Image src='/Logo.svg'></Image>
                </Box>
                <Text
                width='70%'
                fontFamily={fontFamilyNunito}
                fontWeight='700'
                paddingTop={{
                    'sm' : '7%',
                    'lg' : '1%'
                }}
                fontSize={{
                    'sm' : '23px',
                    'lg' : '40px'
                }}
                color='white'
                marginLeft={{
                    'sm' : '5%',
                    'lg' : '5%'
                }}
                alignSelf={{
                    'lg' : 'center'
                }}>My Spend</Text>
            </Flex>
            <Card
            width={{
                'lg' : '50%'
            }}
            borderRadius='14px'
            position='relative'
            marginTop={{
                'lg' : '5%'
            }}
            left={{
                'lg' : '10%'
            }}
            height={{
                'lg' : '85%'
            }}
            marginBottom={{
                'sm' : '10%',
                'lg' : '5%'
            }}>
                <CardHeader>
                    <Text
                    fontFamily={fontFamilyInter}
                    fontSize={{
                        'lg' : '34px'
                    }}>Create Your Account</Text>
                    <Text
                    fontFamily={fontFamilyInter}
                    fontSize={{
                        'lg' : '24px'
                    }}
                    fontWeight='300'>It’s free and easy</Text>
                </CardHeader>
                <CardBody
                alignItems='center'>
                    <Flex
                    flexDir={{
                        'lg' : 'column'
                    }}
                    flexWrap='wrap'
                    gap='1rem'>
                        <Box
                        width='100%'>
                            <FormControl
                            isRequired
                            isInvalid={!isNameValid}>
                                <FormLabel
                                fontFamily={fontFamilyInter}
                                fontSize={{
                                    'lg' : '16px'
                                }}
                                fontWeight='500'>Name</FormLabel>
                                <Input 
                                variant='filled'
                                placeholder='Enter your name'
                                value={name}
                                onChange={nameHandler}></Input>
                                {
                                isNameValid? 
                                <FormHelperText
                                fontFamily={fontFamilyInter}>Format name valid</FormHelperText> 
                                :
                                <FormErrorMessage
                                fontFamily={fontFamilyInter}>Length of name should larger than equals 6</FormErrorMessage>
                                }
                            </FormControl>
                        </Box>
                        <Box
                        width='100%'>
                            <FormControl isRequired isInvalid={!isPasswordValid}>
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
                                {
                                    isPasswordValid? 
                                    <FormHelperText
                                    fontFamily={fontFamilyInter}>Format password valid</FormHelperText>
                                    :
                                    <FormErrorMessage
                                    fontFamily={fontFamilyInter}>Length of password should larger than equals 6</FormErrorMessage>
                                }
                            </FormControl>
                        </Box>
                        <Box
                        width='100%'>
                            <FormControl 
                            isRequired
                            isInvalid={!isEmailValid}>
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
                                {
                                    isEmailValid ? 
                                    <FormHelperText
                                    fontFamily={fontFamilyInter}>Format email valid</FormHelperText>
                                    :
                                    <FormErrorMessage
                                    fontFamily={fontFamilyInter}>Format Email tidak valid</FormErrorMessage>
                                }
                            </FormControl>
                        </Box>
                        <Flex
                        flexDir='row'>
                            <Checkbox
                            onChange={checkHandler}
                            borderColor='black'
                            height={{
                                'lg' : '2%'
                            }}
                            alignSelf={{
                                'lg' : 'center'
                            }}
                            ></Checkbox>
                            <Text
                            fontFamily={fontFamilyInter}
                            marginLeft='10px'>
                            By checking the box, I agree that I have read and accepted the <Link color={blue}>Terms of Use</Link> and <Link color={blue}>Privacy Policy</Link>
                            </Text>
                        </Flex>
                    </Flex>
                </CardBody>
                <CardFooter 
                flexDir='row'>
                    <Box>
                        <Text
                        fontFamily={fontFamilyInter}>Already have an account? <Link 
                            color={blue} 
                            fontFamily={fontFamilyInter}
                            href='/login'>
                                Sign in
                            </Link>
                        </Text>
                    </Box>
                    <Button
                    backgroundColor={blue}
                    colorScheme='blue'
                    marginLeft='30%'
                    onClick={submitHandler}
                    marginTop={{
                        'sm' : '2%',
                        'lg' : '0'
                    }}>
                        <Text
                        fontFamily={fontFamilyInter}
                        fontSize={{
                            'lg' : '14px'
                        }}
                        fontWeight='600'>Create account</Text>
                    </Button>
                </CardFooter>
            </Card>
        </Flex>
    )
}

export default RegisterPage;