# WorldTime SDK utility: make_context

from projectname_sdk.core.context import WorldTimeContext


def make_context_util(ctxmap, basectx):
    return WorldTimeContext(ctxmap, basectx)
