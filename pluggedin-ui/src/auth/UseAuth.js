import React, { createContext, useContext, useEffect, useState, useMemo } from "react";
import eccrypto from "eccrypto"
var CryptoJS = require("crypto-js");

let AuthContext = createContext();


export const AuthProvider = ({children}) => {

    const [user, setUser] = useState(null)
    const [username, setUsername] = useState(null)
    const [error, setError] = useState()
    const [loading, setLoading] = useState()
    const [loadingInitial, setLoadingInitial] = useState(true)
    
    useEffect( () => {
        
        async function validateToken() {
            const jwt = localStorage.getItem("jwt")
            const pk = localStorage.getItem("privatekey")
            const response = await fetch('http://127.0.0.1:1323/api/fluent/username', {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer '+jwt
                }
            })

            if(response.status === 202 && pk !== null) {
                setUser(pk)
                const data = await response.json()
                setUsername(data)
            }
            setLoadingInitial(false)
        }
        validateToken()
    }, [])

    const logout = () => {
        localStorage.clear()
        setUser(null)
    }

    const login = async (username, password) => {
        setLoading(true)
        //make fetch to login user
        const response = await fetch('http://127.0.0.1:1323/api/oauth/signin', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({username, password})
        })

        if(response.status !== 200) {
            setUser(false)
            setLoading(false)
            return false
        } else {
            let responseJson = await response.json()
            localStorage.setItem("jwt", responseJson)
            const key = await fetch('http://127.0.0.1:1323/api/fluent/secret-key', {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer '+responseJson
                }
            })
            const data = await key.json()
            const decrypted = CryptoJS.AES.decrypt(data, password).toString(CryptoJS.enc.Utf8)
            //console.log(decrypted)
            localStorage.setItem("privatekey", decrypted)
            setLoading(false)
            setUser(decrypted)
            return true
        }
    }

    const signup = (username, password) => {
        setLoading(true)

        //make fetch to signup user
    }

    const memoedValue = useMemo(
        () => ({
            loading,
            error,
            logout,
            login,
            signup,
            user,
            username
        }), [loading, error, user, username]
    )
    
    return (
        <AuthContext.Provider value={memoedValue}>
            {!loadingInitial && children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => {
    return useContext(AuthContext)
}