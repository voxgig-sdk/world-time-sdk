<?php
declare(strict_types=1);

// WorldTime SDK utility: result_headers

class WorldTimeResultHeaders
{
    public static function call(WorldTimeContext $ctx): ?WorldTimeResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
