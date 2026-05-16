<?php
declare(strict_types=1);

// WorldTime SDK utility: result_body

class WorldTimeResultBody
{
    public static function call(WorldTimeContext $ctx): ?WorldTimeResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
