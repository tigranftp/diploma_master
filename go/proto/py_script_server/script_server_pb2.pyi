from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class ScriptRequest(_message.Message):
    __slots__ = ("req",)
    REQ_FIELD_NUMBER: _ClassVar[int]
    req: str
    def __init__(self, req: _Optional[str] = ...) -> None: ...

class ScriptResponse(_message.Message):
    __slots__ = ("rsp",)
    RSP_FIELD_NUMBER: _ClassVar[int]
    rsp: str
    def __init__(self, rsp: _Optional[str] = ...) -> None: ...
