/*
 *
 * Copyright 2022 Oleg Borodin  <borodin@unix7.org>
 *
 */


package dsrpc

import (
    "time"
)

func LogRequest(context *Context) error {
    var err error
    logDebug("request:", string(context.reqRPC.JSON()))
    return err
}

func LogResponse(context *Context) error {
    var err error
    logDebug("response:", string(context.resRPC.JSON()))
    return err
}

func LogAccess(context *Context) error {
    var err error
    execTime := time.Now().Sub(context.start)
    logAccess(context.remoteHost, context.reqRPC.Method, execTime)
    return err
}
