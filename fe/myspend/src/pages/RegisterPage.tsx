import { Box, Button, Card, CardBody, CardFooter, CardHeader, Checkbox, Flex, FormControl, FormErrorMessage, FormHelperText, FormLabel, Image, Input, InputGroup, InputRightElement, Link, Text } from "@chakra-ui/react"
import { ChangeEvent, FormEvent, MouseEvent, useState } from "react";
import { blue, fontFamilyInter, fontFamilyNunito } from '../styling/style';

const RegisterPage = () => {

    const expression: RegExp = /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i;

    const [name, setName] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [show, setShow] = useState<boolean>(false);
    const [check, setCheck] = useState<boolean>(false);

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
        setCheck(!check);
    }

    const submitHandler = (e: FormEvent<HTMLButtonElement>) => {
        if(isNameValid && isPasswordValid && isEmailValid && check){
            // call API
        } else {
            e.preventDefault();
        }
    }

    return(
        <Flex
        border='4px solid black'
        width='100%'
        height='100%'
        backgroundImage='/background_auth.svg'
        backgroundSize='cover'
        position='relative'>
            <Flex
            flexWrap='wrap'
            flexDir='row'
            border='4px solid green'
            width={{
                'lg' : '23%'
            }}
            marginLeft={{
                'lg' : '7%'
            }}
            height={{
                'lg' : '100%'
            }}
            alignSelf='flex-end'
            marginBottom={{
                'lg' : '20%'
            }}>
                <Box
                border = '4px solid brown'
                backgroundColor='white'
                borderRadius='14px'>
                    <Image src='/Logo.svg'></Image>
                </Box>
                <Text
                fontFamily={fontFamilyNunito}
                fontWeight='700'
                border='4px solid pink'
                paddingTop={{
                    'lg' : '6%'
                }}
                fontSize={{
                    'lg' : '40px'
                }}
                color='white'>My Spend</Text>
            </Flex>
            <Card
            width={{
                'lg' : '50%'
            }}
            borderRadius='14px'
            border='4px solid yellow'
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
                alignItems='center'
                border='4px solid red'>
                    <Flex
                    flexDir={{
                        'lg' : 'column'
                    }}
                    border='4px solid brown'
                    flexWrap='wrap'
                    gap='1rem'>
                        <Box
                        border='4px solid purple'>
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
                        border='4px solid red'>
                            <FormControl isRequired isInvalid={!isPasswordValid}>
                                <FormLabel
                                fontFamily={fontFamilyInter}
                                fontSize={{
                                    'lg' : '16px'
                                }}
                                fontWeight='500'>Password</FormLabel>
                                <InputGroup
                                border = '4px solid green'>
                                    <Input 
                                    variant='filled'
                                    value={password}
                                    onChange={passwordHandler}
                                    placeholder='Enter your password'
                                    type={show? 'text' : 'password'}></Input>
                                    <InputRightElement
                                    border='4px solid blue'
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
                        border='4px solid blue'>
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
                        border='4px solid green'
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
                border='4px solid grey'
                flexDir='row'>
                    <Box
                    border='4px solid green'>
                        <Text
                        fontFamily={fontFamilyInter}>Already have an account? <Link 
                            color={blue} 
                            fontFamily={fontFamilyInter}>
                                Sign in
                            </Link>
                        </Text>
                    </Box>
                    <Button
                    border='4px solid black'
                    backgroundColor={blue}
                    colorScheme='blue'
                    marginLeft='30%'
                    onSubmit={submitHandler}>
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