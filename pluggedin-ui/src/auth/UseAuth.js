import React, { createContext, useContext, useEffect, useState, useMemo } from "react";
import eccrypto from "eccrypto"

let AuthContext = createContext();


export const AuthProvider = ({children}) => {

    const [user, setUser] = useState(null)
    const [error, setError] = useState()
    const [loading, setLoading] = useState()
    const [loadingInitial, setLoadingInitial] = useState(true)
    
    useEffect( () => {
        
        async function validateToken() {
            const jwt = localStorage.getItem("jwt")
            
            const response = await fetch('http://127.0.0.1:1323/api/fluent/secret-key', {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer '+jwt
                }
            })

            if(response.status !== 202) {
                setUser(null)
                setLoadingInitial(false)
            } else {
                let pk = await response.json()
                console.log(pk)
                if(pk === null) {
                    console.log("BUG llave privada no existe !!")
                }
                setUser(pk)
                setLoadingInitial(false)
            }
        }
        validateToken()
    }, [])

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
            const pk = await fetch('http://127.0.0.1:1323/api/fluent/secret-key', {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer '+responseJson
                }
            })
            const pkJson = await pk.json()
            let privateKey = ""
            console.log(pkJson)
            if(pkJson === null) {
                console.log("no tienes una private key asociada asi que debes generar una")
                privateKey = eccrypto.generatePrivate()
                console.log(privateKey)
                await fetch(`http://127.0.0.1:1323/api/fluent/secret-key`, {
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer '+responseJson
                    },
                    body: JSON.stringify({privateKey})
                })
            }
            
            setLoading(false)
            setUser(privateKey)
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
            login,
            signup,
            user
        }), [loading, error, user]
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