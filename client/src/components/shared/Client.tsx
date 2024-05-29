import { useEffect, useState } from "react";
type ClientProps = React.PropsWithChildren<{}>;

const Client: React.FC<ClientProps> = ({ children }) => {
    const [isLoaded, setIsLoaded] = useState(false)

    useEffect(() => {
      setIsLoaded(true)
    }, [])

  return isLoaded ? children : (
    null
  )
}

export default Client