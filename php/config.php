<?php
declare(strict_types=1);

// WorldTime SDK configuration

class WorldTimeConfig
{
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "WorldTime",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://worldtimeapi.org/api",
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "ipn" => [],
                    "ipn2" => [],
                    "timezone" => [],
                ],
            ],
            "entity" => [
        'ipn' => [
          'fields' => [],
          'name' => 'ipn',
          'op' => [],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'ipn2' => [
          'fields' => [
            [
              'name' => 'abbreviation',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'client_ip',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'datetime',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'day_of_week',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 3,
            ],
            [
              'name' => 'day_of_year',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 4,
            ],
            [
              'name' => 'dst',
              'req' => false,
              'type' => '`$BOOLEAN`',
              'active' => true,
              'index$' => 5,
            ],
            [
              'name' => 'dst_from',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 6,
            ],
            [
              'name' => 'dst_offset',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 7,
            ],
            [
              'name' => 'dst_until',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 8,
            ],
            [
              'name' => 'raw_offset',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 9,
            ],
            [
              'name' => 'timezone',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 10,
            ],
            [
              'name' => 'unixtime',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 11,
            ],
            [
              'name' => 'utc_datetime',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 12,
            ],
            [
              'name' => 'utc_offset',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 13,
            ],
            [
              'name' => 'week_number',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 14,
            ],
          ],
          'name' => 'ipn2',
          'op' => [
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '8.8.8.8',
                        'kind' => 'param',
                        'name' => 'ipv4',
                        'orig' => 'ipv4',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/ip/{ipv4}',
                  'parts' => [
                    'ip',
                    '{ipv4}',
                  ],
                  'select' => [
                    'exist' => [
                      'ipv4',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'method' => 'GET',
                  'orig' => '/ip',
                  'parts' => [
                    'ip',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'args' => [],
                  'select' => [],
                  'index$' => 1,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'ip',
              ],
            ],
          ],
        ],
        'timezone' => [
          'fields' => [
            [
              'name' => 'abbreviation',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 0,
            ],
            [
              'name' => 'client_ip',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 1,
            ],
            [
              'name' => 'datetime',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 2,
            ],
            [
              'name' => 'day_of_week',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 3,
            ],
            [
              'name' => 'day_of_year',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 4,
            ],
            [
              'name' => 'dst',
              'req' => false,
              'type' => '`$BOOLEAN`',
              'active' => true,
              'index$' => 5,
            ],
            [
              'name' => 'dst_from',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 6,
            ],
            [
              'name' => 'dst_offset',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 7,
            ],
            [
              'name' => 'dst_until',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 8,
            ],
            [
              'name' => 'raw_offset',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 9,
            ],
            [
              'name' => 'timezone',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 10,
            ],
            [
              'name' => 'unixtime',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 11,
            ],
            [
              'name' => 'utc_datetime',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 12,
            ],
            [
              'name' => 'utc_offset',
              'req' => false,
              'type' => '`$STRING`',
              'active' => true,
              'index$' => 13,
            ],
            [
              'name' => 'week_number',
              'req' => false,
              'type' => '`$INTEGER`',
              'active' => true,
              'index$' => 14,
            ],
          ],
          'name' => 'timezone',
          'op' => [
            'list' => [
              'name' => 'list',
              'points' => [
                [
                  'method' => 'GET',
                  'orig' => '/timezone',
                  'parts' => [
                    'timezone',
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'args' => [],
                  'select' => [],
                  'index$' => 0,
                ],
              ],
              'input' => 'data',
              'key$' => 'list',
            ],
            'load' => [
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'America',
                        'kind' => 'param',
                        'name' => 'area',
                        'orig' => 'area',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                      [
                        'example' => 'New_York',
                        'kind' => 'param',
                        'name' => 'location',
                        'orig' => 'location',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/timezone/{area}/{location}',
                  'parts' => [
                    'timezone',
                    '{area}',
                    '{location}',
                  ],
                  'select' => [
                    'exist' => [
                      'area',
                      'location',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 0,
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'Europe',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'area',
                        'reqd' => true,
                        'type' => '`$STRING`',
                        'active' => true,
                      ],
                    ],
                  ],
                  'method' => 'GET',
                  'orig' => '/timezone/{area}',
                  'parts' => [
                    'timezone',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'area' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                  'active' => true,
                  'index$' => 1,
                ],
              ],
              'input' => 'data',
              'key$' => 'load',
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'timezone',
              ],
            ],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return WorldTimeFeatures::make_feature($name);
    }
}
